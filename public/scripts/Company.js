import React from 'react'

export default React.createClass({
  render() {
    return <div>Company: {this.props.params.companyId}</div>
  }
})
