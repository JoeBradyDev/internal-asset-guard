import { composePlugins, withNx } from '@nx/next';

const nextConfig = {
  output: 'standalone',

  nx: {
    svgr: false,
  },
};

const plugins = [withNx];
export default composePlugins(...plugins)(nextConfig);
