package com.github.syafdia.sample.androidmvvm

import android.databinding.DataBindingUtil
import android.support.v7.app.AppCompatActivity
import android.os.Bundle
import com.github.syafdia.sample.androidmvvm.databinding.ActivityMainBinding

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        val viewDataBinding = DataBindingUtil
                .setContentView<ActivityMainBinding>(this, R.layout.activity_main)

        viewDataBinding.viewModel = MainViewModel()
    }
}
