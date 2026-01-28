import { Controller, Get, Post, Body, Query } from '@nestjs/common';
import { AssetService } from './asset.service';

@Controller('assets')
export class AssetController {
  constructor(private readonly assetService: AssetService) {}

  @Get()
  async getAssets(@Query('limit') limit = 10, @Query('offset') offset = 0) {
    return this.assetService.findAll(Number(limit), Number(offset));
  }

  @Post()
  async createAsset(@Body() body: any) {
    return this.assetService.create(body);
  }
}
