import React, { Component } from 'react';
import axios from 'axios';

export default class Edit extends Component {
  constructor(props) {
    super(props);
    this.onChangeName = this.onChangeName.bind(this);
    this.onChangeSurname = this.onChangeSurname.bind(this);
    this.onChangeDNI = this.onChangeDNI.bind(this);
    this.onChangeGender = this.onChangeGender.bind(this);
    this.onChangeAddr = this.onChangeAddr.bind(this);
    this.onChangePhone = this.onChangePhone.bind(this);
    this.onChangeWeight = this.onChangeWeight.bind(this);
    this.onSubmit = this.onSubmit.bind(this);

    this.state = {
      dni: '',
      name: '',
      surname: '',
      gender: '',
      addr: '',
      phone:'',
      weight: 0
    }
  }

  componentDidMount() {
    axios.get('http://localhost:4000/user/'+this.props.match.params.id)
      .then(response => {
        this.setState({ 
          name: response.data.name, 
          surname: response.data.surname,
          dni: response.data.dni,
          gender: response.data.gender,
          addr: response.data.addr,
          phone: response.data.phone,
          weight: response.data.weight
        });
      })
      .catch(function (error) {
          console.log(error);
    })
  }

  onChangeName(e) {
    this.setState({
      name: e.target.value
    });
  }
  onChangeSurname(e) {
    this.setState({
      surname: e.target.value
    })  
  }
  onChangeDNI(e) {
    this.setState({
      dni: e.target.value
    })
  }
  onChangeGender(e) {
    this.setState({
      gender: e.target.value
    })
  }
  onChangeAddr(e) {
    this.setState({
      addr: e.target.value
    })
  }
  onChangePhone(e) {
    this.setState({
      phone: e.target.value
    })
  }
  onChangeWeight(e) {
    this.setState({
      weight: parseInt(e.target.value)
    })
  }

  onSubmit(e) {
    e.preventDefault();
    const obj = {
      name: this.state.name,
      surname: this.state.surname,
      dni: this.state.dni,
      gender: this.state.gender,
      addr: this.state.addr,
      phone: this.state.phone,
      weight: this.state.weight
    };
    console.log(obj);
    axios.post('http://localhost:4000/user/'+ this.props.match.params.id, obj)
        .then(res => 
          console.log(res.data),
        );
    
    this.props.history.push('/index');
    window.location.reload()
  }
 
  render() {
    return (
        <div style={{ marginTop: 10 }}>
            <h3 align="center"><b>Update employee information</b></h3>
            <form onSubmit={this.onSubmit}>
                  <div className="form-group">
                      <label>Name:  </label>
                      <input type="text" 
                        className="form-control" 
                        value={this.state.name}
                        onChange={this.onChangeName}
                        />
                  </div>
                  <div className="form-group">
                      <label>Surname: </label>
                      <input type="text" 
                        className="form-control"
                        value={this.state.surname}
                        onChange={this.onChangeSurname}
                        />
                  </div>
                  <div className="form-group">
                      <label>DNI: </label>
                      <input type="text" 
                        className="form-control"
                        value={this.state.dni}
                        onChange={this.onChangeDNI}
                        disabled
                        />
                  </div>
                  <div className="form-group">
                      <label>Gender: </label>
                      <select value={this.state.gender} onChange={this.onChangeGender} className="custom-select">
                        <option value="Male">Male</option>
                        <option value="Female">Female</option>
                      </select>
                  </div>
                  <div className="form-group">
                      <label>Address: </label>
                      <input type="text" 
                        className="form-control"
                        value={this.state.addr}
                        onChange={this.onChangeAddr}
                        />
                  </div>
                  <div className="form-group">
                      <label>Phone: </label>
                      <input type="text" 
                        className="form-control"
                        value={this.state.phone}
                        onChange={this.onChangePhone}
                        />
                  </div>
                  <div className="form-group">
                      <label>Weight: </label>
                      <input type="number" 
                        className="form-control"
                        value={this.state.weight}
                        onChange={this.onChangeWeight}
                        />
                  </div>
                <div className="form-group">
                    <input type="submit" 
                      value="Update Employee" 
                      className="btn btn-primary"/>
                </div>
            </form>
        </div>
    )
  }
}