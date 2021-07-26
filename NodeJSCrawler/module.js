import axios from 'axios';
import cheerio  from 'cheerio';
import FurnitureRepository from './repositories/FurnitureRepository';
import GetFurnitureDetailUseCase from './usecases/GetFurnitureDetailUseCase';
import GetFurniturePreviewsUseCase from './usecases/GetFurniturePreviewsUseCase';

// Repositories.
export const furnitureRepo = new FurnitureRepository(axios, cheerio);

// Use cases.
export const getFurnitureDetailUseCase = new GetFurnitureDetailUseCase(furnitureRepo);
export const getFurniturePreviewsUseCase = new GetFurniturePreviewsUseCase(furnitureRepo);