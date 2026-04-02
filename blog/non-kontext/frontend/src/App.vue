<template>
  <div class="min-h-screen bg-white text-gray-900">
    <AppHeader />
    <main class="max-w-4xl mx-auto px-4 py-8">
      <router-view />
    </main>
    <AppFooter />
  </div>
</template>

<script setup>
import AppHeader from './components/AppHeader.vue'
import AppFooter from './components/AppFooter.vue'
</script>
