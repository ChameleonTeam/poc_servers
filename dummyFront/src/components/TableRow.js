import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import UserProfile from './user.component';

class TableRow extends Component {

  constructor(props) {
        super(props);
        this.delete = this.delete.bind(this);
  }

  delete() {
      axios.delete('http://ec2-34-245-161-251.eu-west-1.compute.amazonaws.com:4000/user/'+this.props.obj.dni)
          .then(console.log('Deleted'), window.location.reload())
          .catch(err => console.log(err))
  }

  render() {
    return (
        <tr>
          <td>
            {this.props.obj.dni}
          </td>
          <td>
            {this.props.obj.name}
          </td>
          <td>
            {this.props.obj.surname}
          </td>
          <td>
            {this.props.obj.gender}
          </td>
          <td>
            {this.props.obj.addr}
          </td>
          <td>
            {this.props.obj.phone}
          </td>
          <td>
            {this.props.obj.weight}
          </td>
          <td hidden={!UserProfile.includeAction("UpdateUser")}>
            <Link to={"/edit/"+this.props.obj.dni} className="btn btn-primary">Edit</Link>
          </td>
          <td hidden={!!UserProfile.includeAction("UpdateUser")}>
            <button className="btn btn-primary" disabled>Edit</button>
          </td>
          <td hidden={!UserProfile.includeAction("DeleteUser")}>
            <button onClick={this.delete} className="btn btn-danger">Delete</button>
          </td>
          <td hidden={!!UserProfile.includeAction("DeleteUser")}>
            <button className="btn btn-danger" disabled>Delete</button>
          </td>
        </tr>
    );
  }
}

export default TableRow;