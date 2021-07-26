import React from 'react';
import axios from 'axios';

class Detail extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isFetchingFurnitureDetail: true,
      furnitureDetail: null,
      errGeneral: '',
    };
  }

  static getInitialProps(ctx) {
    return { productId: Number(ctx.query.id) };
  }

  componentDidMount() {
    this.fetchFurnitureDetail(this.props.productId);
  }

  fetchFurnitureDetail(productId)  {
    this.setState({ isFetchingFurnitureDetail: true });

    axios.get(`/api/furnitures/${productId}`)
      .then((resp) => {
        this.setState({ 
          isFetchingFurnitureDetail: false, 
          furnitureDetail: resp.data,
        });
      })
      .catch((err) => {
        this.setState({ isFetchingFurnitureDetail: false, errGeneral: err });
      })
  }

  render() {
    const { furnitureDetail, isFetchingFurnitureDetail } = this.state;

    if (isFetchingFurnitureDetail) {
      return (
        <section className="page-section">
          <div className="container">
            <div className="text-center">
              <h2 className="section-heading">Fetching Furniture...</h2>
              <h3 className="section-subheading text-muted">
                Please, wait for a moment, while magic is happening
              </h3>
            </div>
          </div>
        </section>
      );
    }

    return (
      <section className="page-section">
        <div className="container">
            <div className="text-center">
                <h2 className="section-heading">Your Furniture</h2>
            </div>
            <br />
            <br />
            <div className="row">
              <div className="col-md-6">
                {furnitureDetail.images.map((v) => (<img src={v} className="img-fluid" />))}
              </div>
              <div className="col-md-6">
                <h4>{furnitureDetail.name}</h4>
                {furnitureDetail.descriptions.map((v) => <p>{v}</p>)}
                <br />
                <h4>Rp. {furnitureDetail.latestPrice.toLocaleString()}</h4>
              </div>
            </div>
        </div>
      </section>
    );
  }
}

export default Detail;