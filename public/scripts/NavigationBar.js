import React from 'react'
import { Link } from 'react-router'

export default React.createClass({
  render() {
    return (
      <ul role="nav" className="navigationBar">
        <Link to="/about" activeClassName="active">About</Link>
        --
        <Link to="/companies" activeClassName="active">Companies</Link>
        --
        <Link to="/company/12" activeClassName="active">Company</Link>
      </ul>
    )
  }
})
