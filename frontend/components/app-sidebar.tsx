'use client';

import * as React from 'react';
import { IconInnerShadowTop, IconListDetails, IconUser } from '@tabler/icons-react';

import { NavMain } from '@/components/nav-main';
import { NavUser } from '@/components/nav-user';
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from '@/components/ui/sidebar';
import { NavDocuments } from './nav-documents';

const data = {
  user: {
    name: 'shadcn',
    email: 'm@example.com',
    avatar: '/avatars/shadcn.jpg',
  },
  navMain: [
    {
      title: 'Vehicles',
      url: '/vehicles',
      icon: IconListDetails,
    },
    {
      title: 'Items',
      url: '/items',
      icon: IconListDetails,
    },
  ],
  navDocuments: [
    {
      name: 'Users',
      url: '/manager/users',
      icon: IconUser,
    },
    {
      name: 'Roles',
      url: '/manager/roles',
      icon: IconUser,
    },
    {
      name: 'Permissions',
      url: '/manager/permission',
      icon: IconUser,
    },
  ],
};

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  return (
    <Sidebar collapsible="offcanvas" {...props}>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton asChild className="data-[slot=sidebar-menu-button]:!p-1.5">
              <a href="#">
                <IconInnerShadowTop className="!size-5" />
                <span className="text-base font-semibold">Acme Inc.</span>
              </a>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
      </SidebarContent>
      <SidebarContent>
        <NavDocuments items={data.navDocuments} />
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={data.user} />
      </SidebarFooter>
    </Sidebar>
  );
}
