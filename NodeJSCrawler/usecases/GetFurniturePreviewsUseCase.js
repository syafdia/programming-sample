export default class GetFurnitureDetailUseCase {

  constructor(furnitureRepo) {
    this.furnitureRepo = furnitureRepo;
  }
  
  // ({ skip: number }) -> Promise<[]FurniturePreview>
  execute(previewUrl, page) {
    return this.furnitureRepo.getFurniturePreviews(previewUrl, page);
  }
}