import React, { Component } from 'react';
import axios from 'axios';
import TableRow from './TableRow';

const endpoint = 'http://localhost:4000/user'

export default class Index extends Component {

  constructor(props) {
      super(props);
      this.state = {persons: []};
    }
    componentDidMount(){
      axios.get(endpoint)
        .then(response => {
          this.setState({ persons: response.data.persons });
        })
        .catch(function (error) {
          console.log(error);
        })
    }
    tabRow(){
      return this.state.persons.map(function(object, i){
          return <TableRow obj={object} key={i} />;
      });
    }

    render() {
      return (
        <div>
          <h3 align="center">Business List</h3>
          <table className="table table-striped" style={{ marginTop: 20 }}>
            <thead>
              <tr>
                <th>DNI</th>
                <th>Name</th>
                <th>Surname</th>
                <th>Gender</th>
                <th>Address</th>
                <th>Phone</th>
                <th>Weight</th>
                <th colSpan="2">Action</th>
              </tr>
            </thead>
            <tbody>
              { this.tabRow() }
            </tbody>
          </table>
        </div>
      );
    }
  }