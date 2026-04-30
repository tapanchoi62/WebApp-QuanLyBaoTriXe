'use client';

import { use } from 'react';
import WarehouseForm from '@/features/inventory/components/WarehouseForm';

export default function Page({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  return <WarehouseForm id={id} />;
}
