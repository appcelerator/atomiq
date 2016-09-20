package stack

import (
	"errors"
	"fmt"
	"path"

	"github.com/appcelerator/amp/api/rpc/service"
	"github.com/appcelerator/amp/data/storage"
	"golang.org/x/net/context"
)

const stackRootKey = "stacks"
const servicesRootKey = "services"
const stackIDLabelName = "io.amp.stack.id"

// Server is used to implement stack.StackService
type Server struct {
	Store storage.Interface
}

// Up implements stack.ServerService Up
func (s *Server) Up(ctx context.Context, in *UpRequest) (*UpReply, error) {
	stack, err := NewStackFromYaml(ctx, in.Stackfile)
	if err != nil {
		return nil, err
	}
	stack.Name = in.StackName
	if err := stackStateMachine.TransitionTo(stack.Id, int32(StackState_Starting)); err != nil {
		return nil, err
	}
	fmt.Printf("Creating stack %s\n", stack.Id)
	s.Store.Create(ctx, path.Join(stackRootKey, stack.Id), stack, nil, 0)
	serviceIDList := make([]string, len(stack.Services), len(stack.Services))
	for i, service := range stack.Services {
		serviceID, err := s.processService(ctx, stack.Id, service)
		if err != nil {
			s.rollbackStack(ctx, stack.Id, serviceIDList, err)
			return nil, err
		}
		serviceIDList[i] = serviceID
	}
	// Save the service id list in ETCD
	val := &ServiceIdList{
		List: serviceIDList,
	}
	fmt.Println("list", val)
	createErr := s.Store.Create(ctx, path.Join(stackRootKey, stack.Id, servicesRootKey), val, nil, 0)
	if createErr != nil {
		s.rollbackStack(ctx, stack.Id, serviceIDList, err)
		return nil, createErr
	}
	if err := stackStateMachine.TransitionTo(stack.Id, int32(StackState_Running)); err != nil {
		return nil, err
	}
	reply := UpReply{
		StackId: stack.Id,
	}
	fmt.Printf("Stack is running: %s\n", stack.Id)
	return &reply, nil
}

// clean up if error happended during stack creation, delete all created services and all etcd data
func (s *Server) rollbackStack(ctx context.Context, stackID string, serviceIDList []string, err error) {
	fmt.Printf("Error clean up stack: %v \n", err)
	server := service.Service{}
	for _, ID := range serviceIDList {
		if ID != "" {
			server.Remove(ctx, ID)
		}
	}
	s.Store.Delete(ctx, path.Join(stackRootKey, stackID), true, nil)
	fmt.Printf("Stack cleaned %s\n", stackID)
}

// start one service and if ok store it in ETCD:
func (s *Server) processService(ctx context.Context, stackID string, serv *service.ServiceSpec) (string, error) {
	serv.Labels[stackIDLabelName] = stackID
	request := &service.ServiceCreateRequest{
		ServiceSpec: serv,
	}
	server := service.Service{}
	reply, err := server.Create(ctx, request)
	if err != nil {
		return "", err
	}
	createErr := s.Store.Create(ctx, path.Join(servicesRootKey, reply.Id), serv, nil, 0)
	if createErr != nil {
		return "", createErr
	}
	return reply.Id, nil
}

// Stop implements stack.ServerService Stop
func (s *Server) Stop(ctx context.Context, in *StackRequest) (*StackReply, error) {
	if running, err := stackStateMachine.Is(in.StackId, int32(StackState_Running)); err != nil {
		return nil, err
	} else if !running {
		return nil, errors.New("Stack is not running")
	}
	fmt.Printf("Stopping stack %s\n", in.StackId)
	server := service.Service{}
	listKeys := &ServiceIdList{}
	err := s.Store.Get(ctx, path.Join(stackRootKey, in.StackId, servicesRootKey), listKeys, true)
	if err != nil {
		return nil, err
	}
	var removeErr error
	for _, key := range listKeys.List {
		err := server.Remove(ctx, key)
		if err != nil {
			removeErr = err
		}

	}
	if removeErr != nil {
		return nil, removeErr
	}
	if err := stackStateMachine.TransitionTo(in.StackId, int32(StackState_Stopped)); err != nil {
		return nil, err
	}
	reply := StackReply{
		StackId: in.StackId,
	}
	fmt.Printf("Stack stopped %s\n", in.StackId)
	return &reply, nil
}

// Remove implements stack.ServerService Remove
func (s *Server) Remove(ctx context.Context, in *StackRequest) (*StackReply, error) {
	if stopped, err := stackStateMachine.Is(in.StackId, int32(StackState_Stopped)); err != nil {
		return nil, err
	} else if !stopped {
		return nil, errors.New("The stack is not stopped")
	}
	fmt.Printf("Removing stack %s\n", in.StackId)
	if err := stackStateMachine.DeleteState(in.StackId); err != nil {
		return nil, err
	}
	if err := s.Store.Delete(ctx, path.Join(stackRootKey, in.StackId, servicesRootKey), true, nil); err != nil {
		return nil, err
	}
	reply := StackReply{
		StackId: in.StackId,
	}
	fmt.Printf("Stack removed %s\n", in.StackId)
	return &reply, nil
}
