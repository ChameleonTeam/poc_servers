import React, { Component } from 'react';
import axios from 'axios';
import UserProfile from './user.component';

const endpoint = 'http://localhost:4000/permissions'

export default class Login extends Component {
  constructor(props) {
      super(props);
      this.onChangeUser = this.onChangeUser.bind(this);
      this.onChangePass = this.onChangePass.bind(this);
      this.onSubmit = this.onSubmit.bind(this);
      this.onLogout = this.onLogout.bind(this);

      this.state = {
          user: '',
          pass: ''
      }
  }
  onChangeUser(e) {
    this.setState({
      user: e.target.value
    });
  }
  onChangePass(e) {
    this.setState({
      pass: e.target.value
    })  
  }

  onSubmit(e) {
    e.preventDefault();
    const obj = {
      user: this.state.user
    };

    axios.post(endpoint, obj)
        .then(res => {
          console.log(res.data);
          if (!!res.data.actions) {
            UserProfile.setName(this.state.user);
            UserProfile.setLogged(true);
            UserProfile.setActions(res.data.actions);
            this.setState({
              user: '',
              pass: ''
            })
          }
        })
        .catch(function (error) {
          console.log(error);
        });
  
    console.log(`The values are ${this.state.user}, ${this.state.pass}`)
  }

  onLogout(e) {
    e.preventDefault();
    UserProfile.setName("");
    UserProfile.setLogged(false);
    UserProfile.setActions([]);
    this.forceUpdate();
  }

  isLogged() {
    return UserProfile.isLogged()
  }
 
  render() {
    console.log("Esta logeado:", this.isLogged());
      return (
        <div>
          <div style={{ marginTop: 10 }} hidden={this.isLogged() === "true"}>
              <h3>Enter your credentials:</h3>
              <form onSubmit={this.onSubmit}>
                  <div className="form-group">
                      <label>User:  </label>
                      <input type="text" 
                        className="form-control" 
                        value={this.state.user}
                        onChange={this.onChangeUser}
                        />
                  </div>
                  <div className="form-group">
                      <label>Pass: </label>
                      <input type="password" 
                        className="form-control"
                        value={this.state.pass}
                        onChange={this.onChangePass}
                        />
                  </div>
                  <div className="form-group">
                      <input type="submit" value="Login" className="btn btn-primary"/>
                  </div>
              </form>
          </div>

          <div style={{ marginTop: 10 }} hidden={this.isLogged() === "false"}>
            <h3>You are logged!</h3>
            <br></br>
            <form onSubmit={this.onLogout}>
              <div className="form-group">
                <input type="submit" value="Logout" className="btn btn-primary"/>
              </div>
            </form>
          </div>
        </div>
      )
  }
}