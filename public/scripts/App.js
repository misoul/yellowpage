import React from 'react'
import NavigationBar from './NavigationBar'

export default React.createClass({
  render() {
    return(
      <div>
        <NavigationBar/>
        {this.props.children}
      </div>
    )
  }
})
