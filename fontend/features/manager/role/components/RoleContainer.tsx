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
import { Checkbox } from '@/components/ui/checkbox';
import http from '@/lib/axios';

export default function RoleContainer() {
  const [roles, setRoles] = useState([]);
  const [permissions, setPermissions] = useState([]);

  const [open, setOpen] = useState(false);
  const [roleName, setRoleName] = useState('');
  const [selected, setSelected] = useState<number[]>([]);
  const [editId, setEditId] = useState<number | null>(null);

  const load = async () => {
    const r = await http.get('/roles');
    const p = await http.get('/permissions');
    setRoles(r.data.data);
    setPermissions(p.data.data);
  };

  useEffect(() => {
    load();
  }, []);

  const togglePermission = (id: number) => {
    setSelected((prev) => (prev.includes(id) ? prev.filter((x) => x !== id) : [...prev, id]));
  };

  const save = async () => {
    if (editId) {
      await http.put(`/roles/${editId}`, {
        name: roleName,
        permissionIds: selected,
      });
    } else {
      await http.post(`/roles`, {
        name: roleName,
        permissionIds: selected,
      });
    }

    setOpen(false);
    setRoleName('');
    setSelected([]);
    setEditId(null);
    load();
  };

  const remove = async (id: number) => {
    await http.delete(`/roles/${id}`);
    load();
  };

  return (
    <div className="p-6 space-y-4">
      <div className="flex justify-between">
        <h1 className="text-xl font-bold">Roles</h1>
        <Button onClick={() => setOpen(true)}>+ Create</Button>
      </div>

      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead>Permissions</TableHead>
            <TableHead className="text-right">Actions</TableHead>
          </TableRow>
        </TableHeader>

        <TableBody>
          {roles.map((r: any, index: number) => (
            <TableRow key={index}>
              <TableCell>{r.Name}</TableCell>
              <TableCell className="max-w-[300px]">
                {r.Permissions?.map((p: any) => (
                  <span
                    key={p.ID}
                    className="inline-block bg-gray-200 px-2 py-1 rounded mr-1 text-xs"
                  >
                    {p.Name}
                  </span>
                ))}
              </TableCell>

              <TableCell className="text-right space-x-2">
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => {
                    setEditId(r.ID);
                    setRoleName(r.Name);
                    setSelected(r.permissions?.map((p: any) => p.ID) || []);
                    setOpen(true);
                  }}
                >
                  Edit
                </Button>

                <Button variant="destructive" size="sm" onClick={() => remove(r.ID)}>
                  Delete
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>

      {/* Modal */}
      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent className="max-w-md">
          <DialogHeader>
            <DialogTitle>{editId ? 'Update Role' : 'Create Role'}</DialogTitle>
          </DialogHeader>

          <div className="space-y-4">
            <Input
              placeholder="Role name"
              value={roleName}
              onChange={(e) => setRoleName(e.target.value)}
            />

            <div className="space-y-2 border p-3 rounded">
              <div className="font-semibold">Permissions</div>
              {permissions.map((p: any) => (
                <div key={p.ID} className="flex items-center gap-2">
                  <Checkbox
                    checked={selected.includes(p.ID)}
                    onCheckedChange={() => togglePermission(p.ID)}
                  />
                  <span>{p.Name}</span>
                </div>
              ))}
            </div>
          </div>

          <DialogFooter>
            <Button onClick={save}>Save</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
