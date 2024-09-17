import { promises as fs } from 'node:fs';
import * as path from 'node:path';

// Definisikan tipe Task
interface Task {
    id: number;
    title: string;
    completed: boolean;
}

// Tentukan path folder dan file
const folderPath = path.join(__dirname, 'task-list');
const filePath = path.join(folderPath, 'task-list.json');

// Fungsi untuk memastikan folder dan file ada, jika tidak akan dibuat
async function ensureFileExists(): Promise<void> {
    try {
        try {
            await fs.access(folderPath);
        } catch {
            await fs.mkdir(folderPath, { recursive: true });
        }

        try {
            await fs.access(filePath);
        } catch {
            const initialData: Task[] = [];
            await fs.writeFile(filePath, JSON.stringify(initialData, null, 2));
        }
    } catch (error) {
        console.error('Terjadi kesalahan saat memastikan file/folder:', error instanceof Error ? error.message : error);
    }
}

// Fungsi untuk membaca daftar tugas dari file
async function readTasks(): Promise<Task[]> {
    try {
        const data = await fs.readFile(filePath, 'utf8');
        return JSON.parse(data) as Task[];
    } catch (error) {
        console.error('Terjadi kesalahan saat membaca file:', error instanceof Error ? error.message : error);
        return [];
    }
}

// Fungsi untuk menambah tugas baru
export async function addTask(title: string): Promise<void> {
    try {
        await ensureFileExists();
        const tasks = await readTasks();

        const taskExists = tasks.some(task => task.title.toLowerCase() === title.toLowerCase());
        if (taskExists) {
            throw new Error(`Tugas dengan judul "${title}" sudah ada dan tidak bisa ditambahkan lagi.`);
        }

        const newTask: Task = {
            id: tasks.length > 0 ? tasks[tasks.length - 1].id + 1 : 1,
            title,
            completed: false
        };

        tasks.push(newTask);
        await fs.writeFile(filePath, JSON.stringify(tasks, null, 2));
    } catch (error) {
        throw error; // Re-throw the error to be handled by the caller
    }
}

// Fungsi untuk melihat daftar tugas
export async function listTasks(): Promise<void> {
    try {
        await ensureFileExists();
        const tasks = await readTasks();
        if (tasks.length === 0) {
            console.log('Tidak ada tugas yang ditemukan.');
            return;
        }

        tasks.forEach(task => {
            console.log(`ID: ${task.id}, Judul: ${task.title}, Selesai: ${task.completed}`);
        });
    } catch (error) {
        throw error;
    }
}

// Fungsi untuk memperbarui status pengerjaan tugas
export async function updateTaskStatus(id: number, completed: boolean): Promise<void> {
    try {
        await ensureFileExists();
        const tasks = await readTasks();

        const taskIndex = tasks.findIndex(task => task.id === id);
        if (taskIndex === -1) {
            throw new Error(`Tugas dengan ID ${id} tidak ditemukan.`);
        }

        tasks[taskIndex].completed = completed;
        await fs.writeFile(filePath, JSON.stringify(tasks, null, 2));
    } catch (error) {
        throw error; // Re-throw the error to be handled by the caller
    }
}

// Fungsi untuk menghapus tugas
export async function deleteTask(id: number): Promise<void> {
    try {
        await ensureFileExists();
        const tasks = await readTasks();

        const taskIndex = tasks.findIndex(task => task.id === id);
        if (taskIndex === -1) {
            throw new Error(`Tugas dengan ID ${id} tidak ditemukan.`);
        }

        tasks.splice(taskIndex, 1);
        await fs.writeFile(filePath, JSON.stringify(tasks, null, 2));
    } catch (error) {
        throw error; // Re-throw the error to be handled by the caller
    }
}
