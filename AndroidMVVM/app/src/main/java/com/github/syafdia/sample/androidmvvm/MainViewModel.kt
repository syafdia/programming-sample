package com.github.syafdia.sample.androidmvvm

import android.arch.lifecycle.ViewModel
import android.databinding.ObservableField
import android.os.Handler

class MainViewModel : ViewModel() {

    val content = ObservableField<String>("Hello")

    init {
        Handler().postDelayed({
            content.set("Data binding awesome !!!")
        }, 3_000)
    }
}