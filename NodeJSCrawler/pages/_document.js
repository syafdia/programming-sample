import Document, { Html, Head, Main, NextScript } from 'next/document'

class BaseDocument extends Document {
  static async getInitialProps(ctx) {
    const initialProps = await Document.getInitialProps(ctx);
    return { ...initialProps };
  }

  render() {
    return (
      <Html>
        <Head>
          <title>Scrap Fabelios</title>
          <link  rel="icon" type="image/x-icon" href="https://m2fabelio.imgix.net/favicon/stores/1/128x128px.png" />
          <link  rel="shortcut icon" type="image/x-icon" href="https://m2fabelio.imgix.net/favicon/stores/1/128x128px.png" />
          <script src="https://use.fontawesome.com/releases/v5.13.0/js/all.js" crossorigin="anonymous"></script>
          <link href="https://fonts.googleapis.com/css?family=Montserrat:400,700" rel="stylesheet" type="text/css" />
          <link href="https://fonts.googleapis.com/css?family=Droid+Serif:400,700,400italic,700italic" rel="stylesheet" type="text/css" />
          <link href="https://fonts.googleapis.com/css?family=Roboto+Slab:400,100,300,700" rel="stylesheet" type="text/css" />
          <link href="/static/css/styles.css" rel="stylesheet" />
        </Head>

        <body>
          <Main />
          <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
          <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/js/bootstrap.bundle.min.js"></script>
          <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery-easing/1.4.1/jquery.easing.min.js"></script>
          <script src="/static/js/scripts.js"></script>
          <NextScript />
        </body>
      </Html>
    )
  }
}

export default BaseDocument;