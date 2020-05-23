export default class GetFurnitureDetailUseCase {
  constructor(furnitureRepo) {
    this.furnitureRepo = furnitureRepo;
  }
  
  // ({ id: number }) -> Promise<FurnitureDetail>
  execute(productId) {
    return this.furnitureRepo.getFurnitureDetail(productId);
  }
}