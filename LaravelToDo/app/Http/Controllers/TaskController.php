<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Task;

class TaskController extends Controller {

    public function findAll() 
    {
       $tasks = Task::orderBy('created_at', 'asc')->get();
       return $this->respSuccess($tasks);
    }

    public function findOne($id) 
    {
        $task = Task::find($id);
        return $this->respSuccess($task);
    }

    public function create(Request $request) 
    {
        try {
            $task = Task::create($request->all());
            return $this->respSuccess($task);
        } catch (\Illuminate\Database\Eloquent\MassAssignmentException $e) {
            return $this->respError("'{$e->getMessage()}' field is not fillable.", 400);
        } catch (\Throwable $e) {
            dd($e);
            return $this->respError($e->getMessage());
        }
    }

    public function update(Request $request, $id) 
    {
        try {
            $ok = Task::findOrFail($id)->update($request->all());
            if ($ok) {
                return $this->respSuccess(null, 200);
            }

            return $this->respError("Failed updating data with ID {$id}.", 500);
        } catch (\Illuminate\Database\Eloquent\ModelNotFoundException $e) {
            return $this->respError($e->getMessage(), 404);
        } catch (\Throwable $e) {
            return $this->respError($e->getMessage());
        }
    }

    public function delete($id) 
    {
        try {
            $ok = Task::findOrFail($id)->delete();
            if ($ok) {
                return $this->respSuccess(null, 204);
            }

            return $this->respError("Failed deleting data with ID {$id}.", 500);
        } catch (\Illuminate\Database\Eloquent\ModelNotFoundException $e) {
            return $this->respError($e->getMessage(), 404);
        } catch (\Throwable $e) {
            return $this->respError($e->getMessage());
        }
    }
}