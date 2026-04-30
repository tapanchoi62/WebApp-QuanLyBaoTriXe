'use client';

import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { loginSchema, registerSchema } from '../models/auth';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { FieldGroup, Field } from '@/components/ui/field';
import http from '@/lib/axios';
import { useRouter } from 'next/navigation';
import Cookies from 'js-cookie';

type Mode = 'login' | 'register';

export default function AuthForm() {
  const [mode, setMode] = useState<Mode>('login');
  const [errorMsg, setErrorMsg] = useState('');
  const router = useRouter();

  const schema = mode === 'login' ? loginSchema : registerSchema;

  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<any>({
    resolver: zodResolver(schema),
  });

  const onSubmit = async (data: any) => {
    try {
      setErrorMsg('');

      if (mode === 'login') {
        const res = await http.post('/api/login', data);

        Cookies.set('token', res.data.token);
        Cookies.set(
          'user',
          JSON.stringify({
            id: res.data.id,
            username: res.data.username,
            role: res.data.role,
            permissions: res.data.permissions ?? [],
          })
        );
        router.push('/dashboard');
      } else {
        const { confirmPassword, ...payload } = data;
        await http.post('/api/register', payload);

        // auto switch sang login
        setMode('login');
        reset();
        alert('Đăng ký thành công!');
      }
    } catch (err: any) {
      setErrorMsg(err?.response?.data?.error || 'Có lỗi xảy ra');
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen">
      <Card className="w-[400px]">
        <CardContent className="p-6">
          <form onSubmit={handleSubmit(onSubmit)}>
            <FieldGroup>
              {/* TITLE */}
              <div className="text-center mb-4">
                <h1 className="text-2xl font-bold">{mode === 'login' ? 'Login' : 'Register'}</h1>
              </div>

              {/* USERNAME */}
              <div>
                <label>Username</label>
                <input {...register('username')} className="border p-2 w-full" />
                {errors.username && <p className="text-red-500">username sai</p>}
              </div>

              {/* PASSWORD */}
              <div className="mt-4">
                <label>Password</label>
                <input type="password" {...register('password')} className="border p-2 w-full" />
                {errors.password && <p className="text-red-500">passwork sai</p>}
              </div>

              {/* CONFIRM PASSWORD (ONLY REGISTER) */}
              {mode === 'register' && (
                <div className="mt-4">
                  <label>Confirm Password</label>
                  <input
                    type="password"
                    {...register('confirmPassword')}
                    className="border p-2 w-full"
                  />
                  {errors.confirmPassword && <p className="text-red-500"></p>}
                </div>
              )}

              {/* ERROR */}
              {errorMsg && <p className="text-red-500">{errorMsg}</p>}

              {/* SUBMIT */}
              <Field>
                <Button type="submit" className="w-full mt-4">
                  {mode === 'login' ? 'Login' : 'Register'}
                </Button>
              </Field>

              {/* TOGGLE */}
              <div className="text-center mt-4 text-sm">
                {mode === 'login' ? (
                  <>
                    Chưa có tài khoản?{' '}
                    <button
                      type="button"
                      className="text-blue-500"
                      onClick={() => {
                        setMode('register');
                        reset();
                      }}
                    >
                      Register
                    </button>
                  </>
                ) : (
                  <>
                    Đã có tài khoản?{' '}
                    <button
                      type="button"
                      className="text-blue-500"
                      onClick={() => {
                        setMode('login');
                        reset();
                      }}
                    >
                      Login
                    </button>
                  </>
                )}
              </div>
            </FieldGroup>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
