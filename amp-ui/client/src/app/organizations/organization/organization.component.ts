import { Component, OnInit, OnDestroy, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';
import { OrganizationsService } from '../../services/organizations.service'
import { ListService } from '../../services/list.service';
import { UsersService } from '../../services/users.service';
import { User } from '../../models/user.model';
import { Organization } from '../../models/organization.model';
import { Team } from '../../models/team.model';
import { ActivatedRoute } from '@angular/router';
import { MenuService } from '../../services/menu.service'

@Component({
  selector: 'app-organization',
  templateUrl: './organization.component.html',
  styleUrls: ['./organization.component.css']


})
export class OrganizationComponent implements OnInit, OnDestroy {
  organization : Organization
  name = ""
  routeSub : any
  modeCreation : boolean = false
  public listUserService : ListService = new ListService()
  public listUserAddedService : ListService = new ListService()
  addedUsers : User[] = []
  users : User[] = []
  initialUserList : User[] = []
  updated = false
  @ViewChild ('f') form: NgForm;

  constructor(
    private route: ActivatedRoute,
    private usersService: UsersService,
    public organizationsService : OrganizationsService,
    private menuService : MenuService) {
    this.listUserService.setFilterFunction(this.match)
    this.listUserAddedService.setFilterFunction(this.match)
  }

  ngOnInit() {
    this.menuService.setItemMenu('organization', 'Edit')
    this.routeSub = this.route.params.subscribe(params => {
      this.name = params['orgName'];
      for (let org of this.organizationsService.organizations) {
        if (org.name == this.name) {
          this.organization = org
        }
      }
      if (this.organization) {
        this.initialUserList = this.usersService.users.slice()
        this.addedUsers = []
        //
        this.users = this.initialUserList.slice()
        this.listUserAddedService.setData(this.addedUsers)
        this.listUserService.setData(this.users)
      }
    })
  }

  ngOnDestroy() {
    this.routeSub.unsubscribe();
  }

  match(item : User, value : string) : boolean {
    if (value == '' || item.name==='') {
      return true
    }
    if (item.name && item.name.includes(value)) {
      return true
    }
    return false
  }

  addUser( user : User) {
    let list : User[] = []
    for (let item of this.users) {
      if (item.name !== user.name) {
        list.push(item)
      }
    }
    this.users=list
    this.listUserService.setData(this.users)
    this.addedUsers.push(user)
    this.listUserAddedService.setData(this.addedUsers)
    this.updated=true
  }

  removeUser( user : User) {
    let list : User[] = []
    for (let item of this.addedUsers) {
      if (item.name !== user.name) {
        list.push(item)
      }
    }
    this.addedUsers=list
    this.listUserAddedService.setData(this.addedUsers)
    this.users.push(user)
    this.listUserService.setData(this.users)
    if (this.users.length === this.initialUserList.length) {
      this.updated=false
    }
  }

  addAll() {
    this.addedUsers=this.initialUserList.slice()
    this.users=[]
    this.listUserAddedService.setData(this.addedUsers)
    this.listUserService.setData(this.users)
    this.updated=true
  }

  removeAll() {
    this.updated=false
    this.users = this.initialUserList.slice()
    this.addedUsers=[]
    this.listUserAddedService.setData(this.addedUsers)
    this.listUserService.setData(this.users)
  }

  applyUsers() {
    //apply
    this.organization.members = this.addedUsers.slice()
    this.updated=false
  }

}