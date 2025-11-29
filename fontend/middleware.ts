import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export function middleware(req: NextRequest) {
  const token = req.cookies.get("token")?.value;

  const url = req.nextUrl.clone();

  // Danh sách page cần auth
  const protectedPaths = ["/dashboard", "/maintenance", "/inventory"];

  if (protectedPaths.some((path) => url.pathname.startsWith(path))) {
    if (!token) {
      // Chưa login → redirect về login
      url.pathname = "/login";
      return NextResponse.redirect(url);
    }
    // TODO: decode JWT và kiểm tra expiry nếu muốn
  }

  return NextResponse.next();
}

// Áp dụng cho các route cần middleware
export const config = {
  matcher: ["/dashboard/:path*", "/maintenance/:path*", "/inventory/:path*"],
};
