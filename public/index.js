import React from 'react'
import { render } from 'react-dom'
import { Router, Route, hashHistory } from 'react-router'
import NavigationBar from './scripts/NavigationBar'
import About from './scripts/About'
import App from './scripts/App'
import Company from './scripts/Company'
import Companies from './scripts/Companies'

render((
  <Router history={hashHistory}>
    <Route path="/" component={App}>
      <Route path="/about" component={About}/>
      <Route path="/companies" component={Companies}/>
      <Route path="/company/:companyId" component={Company}/>
    </Route>
  </Router>
), document.getElementById('content'))
