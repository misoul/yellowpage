import React from 'react'
import { render } from 'react-dom'
import { Router, Route, hashHistory } from 'react-router'
import App from './scripts/App'
import About from './scripts/About'
import Repos from './scripts/Repos'
import companies from './scripts/companies'

render((
  <Router history={hashHistory}>
    <Route path="/" component={App}/>
    <Route path="/repos" component={Repos}/>
    <Route path="/about" component={companies}/>
  </Router>
), document.getElementById('content'))
