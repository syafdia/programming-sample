<?php

use Illuminate\Database\Seeder;
use App\Models\Task;

class TasksTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        Task::truncate();

        Task::create([
            'name' => 'Create article',
            'status' => Task::STATUS_ACTIVE, 
        ]);

        Task::create([
            'name' => 'Learn cooking',
            'status' => Task::STATUS_ACTIVE, 
        ]);

        Task::create([
            'name' => 'One week one book',
            'status' => Task::STATUS_COMPLETED, 
        ]);
    }
}
