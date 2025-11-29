"use client";

import * as React from "react";
import { Button, } from "@/components/ui/button";
import { ConfirmDialog } from "@/components/ConfirmDialog";
type ButtonProps = React.ComponentProps<typeof Button>;

interface ConfirmButtonProps extends ButtonProps {
  dialogTitle?: string;
  dialogDescription?: string;
  confirmText?: string;
  cancelText?: string;
  onConfirm: () => void;
  children: React.ReactNode;
}

export function ConfirmButton({
  dialogTitle = "Are you sure?",
  dialogDescription = "This action cannot be undone.",
  confirmText = "Confirm",
  cancelText = "Cancel",
  onConfirm,
  children,
  ...buttonProps
}: ConfirmButtonProps) {
  const [open, setOpen] = React.useState(false);

  return (
    <>
      <Button {...buttonProps} onClick={() => setOpen(true)}>
        {children}
      </Button>

      <ConfirmDialog
        open={open}
        onOpenChange={setOpen}
        title={dialogTitle}
        description={dialogDescription}
        confirmText={confirmText}
        cancelText={cancelText}
        onConfirm={onConfirm}
      />
    </>
  );
}

