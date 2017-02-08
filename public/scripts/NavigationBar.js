import React from 'react'
import { Link } from 'react-router'

export default React.createClass({
  render() {
    return (
      <nav role="nav" className="navigationBar navbar navbar-default navbar-static-top">
        <div className="container-fluid">
          <div className="navbar-header">
            <a className="navbar-brand" href="#">YellowPage</a>
          </div>
          <ul className="nav navbar-nav">
            <li><Link to="/about" activeClassName="active">About</Link></li>
            <li><Link to="/companies" activeClassName="active">Companies</Link></li>
            <li><Link to="/company/12" activeClassName="active">Company</Link></li>
            <li><Link to="/comments" activeClassName="active">Comments</Link></li>
          </ul>
          <ul className="nav navbar-nav navbar-right">
            <li><a href="#/signup"><span className="glyphicon glyphicon-user"></span> Sign Up</a></li>
            <li><a href="#/login"><span className="glyphicon glyphicon-log-in"></span> Login</a></li>
          </ul>
        </div>
      </nav>
    )
  }
})
