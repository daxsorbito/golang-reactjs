import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  componentWillMount() {
    const apiPath = process.env.REACT_APP_PUBLIC_URL || "."
    console.log('APIPATH>>>', `${apiPath}/api`)
    fetch(`${apiPath}/api`, {
      mode: 'cors',
      // headers: {
      //   'Accept': 'application/json',
      //   'Content-Type': 'application/json'
      // }
    }).then(results => {
      // console.log('Results>>>', results.json())
      // console.log('results>>>', results.body)
      return results.json()
    }).then(re => {
      console.log('callssss>>', re)
    }).catch(e => console.log('eerrro>>>', e))
  }
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React testing</h1>
        </header>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
      </div>
    );
  }
}

export default App;
