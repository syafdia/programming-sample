import { getFurnitureDetailUseCase } from './../../../module';

export default function (req, res) {
    const {
      query: { id },
      method,
    } = req;
  
    switch (method) {
      case 'GET':
        getFurnitureDetailUseCase
          .execute(id)
          .then((data) => res.status(200).json(data));
        break
      default:
        res.setHeader('Allow', ['GET'])
        res.status(405).end(`Method ${method} Not Allowed`)
    }
  }