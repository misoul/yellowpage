// 'use strict';

var ENTER_KEY = 13;
var currentSearch = "";

var Company = React.createClass({
  rawMarkup: function(str) {
    var md = new Remarkable();
    var rawMarkup = md.render(str);
    return { __html: rawMarkup };
  },

  render: function() {
    var c = this.props.company
    return (
      <div className="company">
        <h4 className="companySummary"> {c.name} </h4>
        <span dangerouslySetInnerHTML={this.rawMarkup(c.desc.toString())} />
        <table><tbody>
          <tr><th>Industry</th><td>{c.industries}</td></tr>
          <tr><th>Website</th><td>{c.website}</td></tr>
          <tr><th>Founding</th><td>{c.foundDate}</td></tr>
          <tr><th>Website</th><td>{c.website}</td></tr>
          <tr><th>Stock</th><td>{c.stockCode}</td></tr>
          <tr><th>About</th><td>{c.desc}</td></tr>
        </tbody></table>
      </div>
    );
  }
});

var CompanyBox = React.createClass({
  loadDataFromServer: function(searchWords = "") {
    $.ajax({
      url: this.props.url,
      data: {keywords: searchWords},
      dataType: 'json',
      cache: false,
      success: function(data) {
        this.setState({data: data});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error(this.props.url, status, err.toString());
      }.bind(this)
    });
  },
  handleSearchChange: function (event) { //TODO: search only when typing stops
    currentSearch = event.target.value.trim();
    this.loadDataFromServer(currentSearch);
  },
  handleNewTodoKeyDown: function (event) {
    if (event.keyCode !== ENTER_KEY) {
      return;
    }
    event.preventDefault();
    this.loadDataFromServer();
  },
  getInitialState: function() {
    return {data: [], currentSearch: ""};
  },
  componentDidMount: function() {
    this.loadDataFromServer();
    // setInterval(this.loadDataFromServer, this.props.pollInterval);
  },
  render: function() {
    return (
      <div className="companyBox">
        <input
          className="searchBox"
          placeholder="Search company.."
          value={currentSearch}
          onKeyDown={this.handleNewTodoKeyDown}
          onChange={this.handleSearchChange}
          autoFocus={true}
        />
        <CompanyList data={this.state.data} />
      </div>
    );
  }
});

var CompanyList = React.createClass({
  render: function() {
    console.log("Rendering: " + this.props.data)
    var companyNodes = this.props.data.map(function(company) {
      return (
        <Company key={company.id} company={company}/>
      );
    });

    return (
      <div className="companyList">
        {companyNodes}
      </div>
    );
  }
});


ReactDOM.render(
  <CompanyBox url="http://localhost:8080/companies" pollInterval={2000} />,
  document.getElementById('content')
);
