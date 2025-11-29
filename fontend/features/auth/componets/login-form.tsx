"use client"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import {
  Field,
  FieldDescription,
  FieldGroup,
} from "@/components/ui/field"
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react"
import http from "@/lib/axios"
import { LoginInput, loginSchema } from "../models/auth"
import { useRouter } from "next/navigation"
import Cookies from "js-cookie";

export default function LoginForm({
  className,
  ...props
}: React.ComponentProps<"div">) {
  const router = useRouter();
  const [errorMsg, setErrorMsg] = useState("");
  const {
    register,
    formState: { errors },
    handleSubmit
  } = useForm<LoginInput>({
    resolver: zodResolver(loginSchema),
  });

  const onSubmit = async (data: LoginInput) => {
    try {
      setErrorMsg("");
      const res = await http.post("/login", data);
      Cookies.set("token", res.data.token, { expires: 1 }); // 1 ngày

      
      router.push("/dashboard");
      alert("Login thành công!");
    } catch (err: any) {
      setErrorMsg(err.response?.data?.error || "Login thất bại");
    }
  };


  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card className="overflow-hidden p-0">
        <CardContent className="grid p-0 md:grid-cols-2">
          <form className="p-6 md:p-8" onSubmit={handleSubmit(onSubmit)}>
            <FieldGroup>
              <div className="flex flex-col items-center gap-2 text-center">
                <h1 className="text-2xl font-bold">Welcome back</h1>
                <p className="text-muted-foreground text-balance">
                  Login to your Acme Inc account
                </p>
              </div>
              <div>
              <label>Username</label>
                <input {...register("username")} className="border p-2 w-full" />
                {errors.username && <p className="text-red-500">{errors.username.message}</p>}
              </div>

              <div className="mt-4">
                <label>Password</label>
                <input type="password" {...register("password")} className="border p-2 w-full" />
                {errors.password && <p className="text-red-500">{errors.password.message}</p>}
              </div>

              <Field>
                <Button type="submit">Login</Button>
              </Field>
            </FieldGroup>
          </form>
          <div className="bg-muted relative hidden md:block">
            <img
              src="/globe.svg"
              alt="Image"
              className="absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
            />
          </div>
        </CardContent>
      </Card>
      <FieldDescription className="px-6 text-center">
        By clicking continue, you agree to our <a href="#">Terms of Service</a>{" "}
        and <a href="#">Privacy Policy</a>.
      </FieldDescription>
    </div>
  )
}
