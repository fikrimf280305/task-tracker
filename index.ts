import yargs from 'yargs';
import { hideBin } from 'yargs/helpers';
import { addTask, listTasks, updateTaskStatus, deleteTask } from './taskManager'; // Impor fungsi dari file taskManager

// Definisikan antarmuka untuk argumen
interface Arguments {
    title?: string;
    id?: number;
    completed?: boolean;
}

// Konfigurasi yargs
yargs(hideBin(process.argv))
    .command<Arguments>(
        'add-tugas <title>',
        'Tambahkan tugas baru',
        (yargs) => {
            return yargs.positional('title', {
                describe: 'Judul tugas yang akan ditambahkan',
                type: 'string'
            });
        },
        async (argv) => {
            try {
                await addTask(argv.title!);
                console.log(`Tugas "${argv.title}" telah berhasil ditambahkan.`);
            } catch (error) {
                console.error('Terjadi kesalahan saat menambahkan tugas:', error instanceof Error ? error.message : error);
            }
        }
    )
    .command(
        'lihat-tugas',
        'Lihat daftar tugas',
        () => {},
        async () => {
            try {
                await listTasks();
            } catch (error) {
                console.error('Terjadi kesalahan saat melihat tugas:', error instanceof Error ? error.message : error);
            }
        }
    )
    .command<Arguments>(
        'update-tugas <id> <completed>',
        'Perbarui status pengerjaan tugas',
        (yargs) => {
            return yargs
                .positional('id', {
                    describe: 'ID tugas yang akan diperbarui',
                    type: 'number'
                })
                .positional('completed', {
                    describe: 'Status pengerjaan tugas (true/false)',
                    type: 'boolean'
                });
        },
        async (argv) => {
            try {
                if (argv.id === undefined || argv.completed === undefined) {
                    throw new Error('ID dan status pengerjaan harus disediakan.');
                }
                await updateTaskStatus(argv.id, argv.completed);
                console.log(`Status tugas dengan ID ${argv.id} telah diperbarui.`);
            } catch (error) {
                console.error('Terjadi kesalahan saat memperbarui status tugas:', error instanceof Error ? error.message : error);
            }
        }
    )
    .command<Arguments>(
        'hapus-tugas <id>',
        'Hapus tugas berdasarkan ID',
        (yargs) => {
            return yargs.positional('id', {
                describe: 'ID tugas yang akan dihapus',
                type: 'number'
            });
        },
        async (argv) => {
            try {
                if (argv.id === undefined) {
                    throw new Error('ID tugas harus disediakan.');
                }
                await deleteTask(argv.id);
                console.log(`Tugas dengan ID ${argv.id} telah berhasil dihapus.`);
            } catch (error) {
                console.error('Terjadi kesalahan saat menghapus tugas:', error instanceof Error ? error.message : error);
            }
        }
    )
    .help()
    .argv; // Parse argumen dan ambil hasil
