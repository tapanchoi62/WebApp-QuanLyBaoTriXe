'use client';

import UserForm from '@/features/users/componnets/UserForm';
import { useParams } from 'next/navigation';

export default function Page() {
  const params = useParams();
  const idParam = Array.isArray(params.id) ? params.id[0] : params.id;

  return <UserForm id={idParam} />;
}
