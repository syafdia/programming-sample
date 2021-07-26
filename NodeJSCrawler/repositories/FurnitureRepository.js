import FurnitureDetail from './../entities/FurnitureDetail';
import FurniturePreview from './../entities/FurniturePreview';

export default class FurnitureRepository {
  constructor(axios, cheerio) {
    this.axios = axios;
    this.cheerio = cheerio;
    this.baseUrl = 'https://fabelio.com';
  }

  getFurnitureDetail(productId) {
    return this.axios
      .get(`${this.baseUrl}/insider/data/product/id/${productId}`)
      .then((resp) => {
        if (resp.status !== 200 || !resp.data.product) {
          return Promise.reject("error fetching furniture detail");
        }

        const product = resp.data.product;

        return this.axios
          .get(decodeURI(product.url))
          .then((resp) => ({ resp, product }))
      })
      .then(({ resp, product }) => {
        if (resp.status !== 200) {
          return Promise.reject("error fetching furniture description");
        }

        return new FurnitureDetail(
          product.name,
          Number(product.unit_price),
          this.parseHtmlToProductDescriptions(resp.data),
          [decodeURI(product.product_image_url)],
        )
      });
  }

  getFurniturePreviews(previewUrl, page) {
    const previewUrlWithPage = `${previewUrl}?p=${page}&ajax=1`;

    return this.axios
      .get(previewUrlWithPage)
      .then((resp) => {
        if (resp.status !== 200 || !resp.data.success) {
          return Promise.reject("error fetching furniture previews");
        }

        return this.parseHtmlToFurniturePreviews(resp.data.html.products_list);
      });
  }

  parseHtmlToProductDescriptions(html) {
    const $ = this.cheerio.load(html);
    const descriptions = $('.product-info__description')
      .find('p')
      .map((i, el) => $(el).text())
      .toArray();

    return descriptions;
  }

  parseHtmlToFurniturePreviews(html) {
    const $ = this.cheerio.load(unescape(html));

    return $('.product-item-info')
      .find('.product-item-details')
      .map((i, el) => {
        const $el = $(el);
        const name = $el.find('.product-item-name > a').attr('title').trim();
        const latestPrice = Number($el.find('.price-container > span').data('price-amount'));
        const detailUrl = $el.find('.product-item-name > a').attr('href').trim();
        const productId = Number($el.find('.price-final_price').data('product-id'));

        return new FurniturePreview(productId, name, latestPrice, detailUrl);
      })
      .toArray();
  }

}