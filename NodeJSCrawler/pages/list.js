import React from 'react';
import Router from 'next/router';
import axios from 'axios';
import Link from 'next/link';


class List extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      errGeneral: '',
      furniturePreviewUrl: '',
      isFetchingfurniturePreviews: false,
      page: 1,
      furniturePreviews: [],
    };
  }

  static getInitialProps(ctx) {
    return { furniturePreviewUrl: ctx.query.url };
  }

  componentDidMount() {
    this.fetchFurniturePreviews();
  }

  fetchFurniturePreviews() {
    const { furniturePreviewUrl } = this.props;
    const { page, furniturePreviews } = this.state;
    
    this.setState({ isFetchingfurniturePreviews: true });

    axios.get(`/api/furnitures?url=${furniturePreviewUrl}&page=${page}`)
      .then((resp) => {
        this.setState({ 
          isFetchingfurniturePreviews: false, 
          furniturePreviews: furniturePreviews.concat(resp.data),
        });
      })
      .catch((err) => {
        this.setState({ isFetchingfurniturePreviews: false, errGeneral: err });
      })
  }

  onClickLoadMore() {
    return (e) => {
      e.preventDefault();
      const { isFetchingfurniturePreviews, page } = this.state;
      if (isFetchingfurniturePreviews) {
        return;
      }

      this.setState({ page: page + 1 }, () => this.fetchFurniturePreviews());
    }
  }

  render() {
    const { errGeneral, furniturePreviews, isFetchingfurniturePreviews } = this.state;
    return (
      <section className="page-section">
        <div className="container">
            <div className="text-center">
                <h2 className="section-heading">Available Furnitures</h2>
                <h3 className="section-subheading text-muted">List of all available furnitures</h3>
            </div>
            <table className="table">
              <thead>
                <tr>
                  <th scope="col">#</th>
                  <th scope="col">Name</th>
                  <th scope="col">Latest Price</th>
                  <th scope="col">URL</th>
                  <th scope="col">Action</th>
                </tr>
              </thead>
              <tbody>
                {furniturePreviews.map((v, i) => (
                  <tr>
                    <th scope="row">{i + 1}</th>
                    <td>{v.name}</td>
                    <td>Rp. {v.latestPrice.toLocaleString()}</td>
                    <td>
                      <a href={v.detailUrl} target="_blank">{v.detailUrl}</a>
                    </td>
                    <td>
                      <Link href={`/detail?id=${v.productId}`}>
                        <a>Detail</a>
                      </Link>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
            <br />
            <div className="text-center">
              <button className="btn btn-primary btn-lg" onClick={this.onClickLoadMore()}>
                {isFetchingfurniturePreviews ? 'Fetching...' : 'Load more'}
              </button>
            </div>
            <br />
        </div>
      </section>
    );
  }
}

export default List;