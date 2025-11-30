'use client';

import { useEffect, useState } from 'react';
import Cookies from 'js-cookie';
export function usePermission(required: string | string[]) {
  const [allowed, setAllowed] = useState(false);

  useEffect(() => {
    const user = JSON.parse(Cookies.get('user') || '{}');

    const userPermissions: string[] = user.permissions || [];

    if (Array.isArray(required)) {
      setAllowed(required.some((p) => userPermissions.includes(p)));
    } else {
      setAllowed(userPermissions.includes(required));
    }
  }, [required]);

  return allowed;
}
