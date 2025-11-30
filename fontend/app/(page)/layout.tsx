import type { ReactNode } from 'react';
import { SidebarInset, SidebarProvider } from '@/components/ui/sidebar';
import { AppSidebar } from '@/components/app-sidebar';
import { HeaderProvider } from '@/contexts/HeaderContext';
import Header from '@/components/Header';
import { Toaster } from '@/components/ui/sonner';

export default function Layout({ children }: { children: ReactNode }) {
  return (
    <div>
      <SidebarProvider
        style={
          {
            '--sidebar-width': 'calc(var(--spacing) * 72)',
            '--header-height': 'calc(var(--spacing) * 12)',
          } as React.CSSProperties
        }
      >
        <AppSidebar variant="inset" />
        <SidebarInset>
          <div className="flex flex-1 flex-col">
            <HeaderProvider>
              <Header />
              {children}
            </HeaderProvider>
          </div>
        </SidebarInset>
      </SidebarProvider>
      <Toaster />
    </div>
  );
}
