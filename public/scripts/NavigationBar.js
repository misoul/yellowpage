import React from 'react'
import { Link } from 'react-router'

export default React.createClass({
  render() {
    return (
      <div className="navigationBar">
        <Link to="/about">About</Link>
        --
        <Link to="/repos">Repos</Link>
      </div>
    )
  }
})
