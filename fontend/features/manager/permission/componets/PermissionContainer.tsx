'use client';

import { useEffect, useState } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import {
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
} from '@/components/ui/table';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from '@/components/ui/dialog';
import http from '@/lib/axios';

export default function PermissionContainer() {
  const [permissions, setPermissions] = useState([]);
  const [open, setOpen] = useState(false);
  const [formName, setFormName] = useState('');
  const [editId, setEditId] = useState<number | null>(null);

  const load = async () => {
    const res = await http.get('/permissions');
    setPermissions(res.data.data);
  };

  useEffect(() => {
    load();
  }, []);

  const save = async () => {
    if (editId) {
      await http.put(`/permissions/${editId}`, { name: formName });
    } else {
      await http.post('/permissions', { name: formName });
    }
    setOpen(false);
    setFormName('');
    setEditId(null);
    load();
  };

  const remove = async (id: number) => {
    await http.delete(`/permissions/${id}`);
    load();
  };

  return (
    <div className="p-6 space-y-4">
      <div className="flex justify-between">
        <h1 className="text-xl font-bold">Permissions</h1>
        <Button onClick={() => setOpen(true)}>+ Create</Button>
      </div>

      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead className="text-right">Actions</TableHead>
          </TableRow>
        </TableHeader>

        <TableBody>
          {permissions.map((p: any, index: number) => (
            <TableRow key={index}>
              <TableCell>{p.Name}</TableCell>
              <TableCell className="text-right space-x-2">
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => {
                    setEditId(p.ID);
                    setFormName(p.Name);
                    setOpen(true);
                  }}
                >
                  Edit
                </Button>
                <Button variant="destructive" size="sm" onClick={() => remove(p.ID)}>
                  Delete
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>

      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{editId ? 'Update Permission' : 'Create Permission'}</DialogTitle>
          </DialogHeader>

          <Input
            placeholder="Permission name"
            value={formName}
            onChange={(e) => setFormName(e.target.value)}
          />

          <DialogFooter>
            <Button onClick={save}>Save</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
