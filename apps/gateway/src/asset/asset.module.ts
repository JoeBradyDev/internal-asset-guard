import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { AssetController } from './asset.controller';
import { AssetService } from './asset.service';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'ASSET_PACKAGE',
        transport: Transport.GRPC,
        options: {
          package: 'asset_service',
          // Note: We'll adjust project.json to copy into 'asset' instead of 'assets'
          protoPath: join(__dirname, 'asset/asset_service.proto'),
          url: 'asset-service:50051',
        },
      },
    ]),
  ],
  controllers: [AssetController],
  providers: [AssetService],
})
export class AssetModule {}
