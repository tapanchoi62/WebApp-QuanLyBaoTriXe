import { z } from 'zod';

export const loginSchema = z.object({
  username: z.string().min(3, 'Username phải từ 3 ký tự trở lên'),
  password: z.string().min(6, 'Password phải từ 6 ký tự trở lên'),
});

export const registerSchema = loginSchema
  .extend({
    confirmPassword: z.string(),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: 'Password không khớp',
    path: ['confirmPassword'],
  });

export type LoginInput = z.infer<typeof loginSchema>;
export type RegisterInput = z.infer<typeof registerSchema>;
