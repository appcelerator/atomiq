#!/bin/bash

test_setup() {
  amp -k user signup --name user1 --password password --email email@user1.amp
}

test_name() {
  find $HOME/.config/amp -name 'localhost*'
}

test_teardown() {
  amp -k user rm user1
}
