import React from 'react';
import Link from 'next/link';
import Router from 'next/router';

class Index extends React.Component {
  constructor(props) {
    super(props);
    this.validSearchUrlPrefix = 'https://fabelio.com/cp';
    this.state = {
      searchUrl: '',
      errGeneral: '',
    };
  }

  onChangeSearchUrl() {
    return (e) => {
      e.preventDefault();
      this.setState({ errGeneral: '', searchUrl: e.target.value.trim() });
    }
  }

  onClickSearch() {
    return (e) => {
      e.preventDefault();

      if (this.state.searchUrl === '') {
        this.setState({ errGeneral: 'URL is not valid.' })
        return;
      }

      if (!this.state.searchUrl.includes(this.validSearchUrlPrefix)) {
        this.setState({ errGeneral: 'URL is not valid.' })
        return;
      }

      Router.push({
        pathname: '/list',
        query: { url: this.state.searchUrl },
      });
    }
  }

  render() {
    return (
      <section className="page-section">
        <div className="container">
          <div className="text-center">
            <h2 className="section-heading">Search Furniture</h2>
            <h3 className="section-subheading text-muted">Search all furnitures by category URL</h3>
          </div>
          <div className="row">
            <div className="col-md-12">
              <div className="form-group">
                  <input 
                    className="form-control" 
                    type="text" 
                    placeholder={`Eg: ${this.validSearchUrlPrefix}/your-dream-furniture.html`}
                    onChange={this.onChangeSearchUrl()}
                  />
                  <p className="help-block text-danger">&nbsp;{this.state.errGeneral}</p>
              </div>
            </div>
          </div>
          <br />
          <div className="text-center">
            <button className="btn btn-primary btn-lg" onClick={this.onClickSearch()}>
              Start Search
            </button>
          </div>
          <br />
        </div>
      </section>
    );
  }
}

export default Index;