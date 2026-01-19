import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export function middleware(request: NextRequest) {
  // Check if the request is for the API
  if (request.nextUrl.pathname.startsWith('/api')) {
    const gatewayUrl = process.env.GATEWAY_URL || 'http://localhost:3000';

    // Construct the new URL for the Gateway
    const targetUrl = new URL(request.nextUrl.pathname, gatewayUrl);

    return NextResponse.rewrite(targetUrl);
  }
}

export const config = {
  matcher: '/api/:path*',
};
