import React from 'react'
import { render } from 'react-dom'
import { Router, Route, hashHistory } from 'react-router'
import NavigationBar from './scripts/NavigationBar'
import About from './scripts/About'
import Repos from './scripts/Repos'
import Company from './scripts/Company'
import Companies from './scripts/Companies'

render((
  <Router history={hashHistory}>
    <Route path="/" component={Companies}/>
    <Route path="/repos" component={Repos}/>
    <Route path="/about" component={About}/>
    <Route path="/companies" component={Companies}>
      <Route path="/companies/:companyId" component={Company}
    </Route>
  </Router>
), document.getElementById('content'))
