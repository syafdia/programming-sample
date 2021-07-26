import { getFurniturePreviewsUseCase } from './../../module';

export default function (req, res) {
  getFurniturePreviewsUseCase
    .execute(req.query.url, req.query.page || 1)
    .then((data) => res.status(200).json(data));
}