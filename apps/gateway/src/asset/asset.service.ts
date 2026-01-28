import { Injectable, OnModuleInit, Inject } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { firstValueFrom } from 'rxjs';

interface AssetGrpcService {
  listAssets(data: { limit: number; offset: number }): any;
  createAsset(data: any): any;
}

@Injectable()
export class AssetService implements OnModuleInit {
  private assetGrpc: AssetGrpcService;

  constructor(@Inject('ASSET_PACKAGE') private client: ClientGrpc) {}

  onModuleInit() {
    this.assetGrpc = this.client.getService<AssetGrpcService>('AssetService');
  }

  async findAll(limit: number, offset: number) {
    return firstValueFrom(this.assetGrpc.listAssets({ limit, offset }));
  }

  async create(data: any) {
    return firstValueFrom(this.assetGrpc.createAsset(data));
  }
}
