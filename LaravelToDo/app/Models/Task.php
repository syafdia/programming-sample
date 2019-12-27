<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Task extends Model 
{
    const STATUS_ACTIVE = 1;
    const STATUS_COMPLETED = 2;

    protected $fillable = ['name', 'status'];
}