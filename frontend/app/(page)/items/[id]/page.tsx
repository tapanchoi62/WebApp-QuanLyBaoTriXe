'use client';

import ItemForm from '@/features/items/components/ItemForm';
import { useParams } from 'next/navigation';

export default function Page() {
  const params = useParams();
  const idParam = Array.isArray(params.id) ? params.id[0] : params.id;

  return <ItemForm id={idParam} />;
}
