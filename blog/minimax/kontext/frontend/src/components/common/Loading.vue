<script setup lang="ts">
defineProps<{
  loading?: boolean
}>()
</script>

<template>
  <div v-if="loading" class="flex justify-center items-center py-20">
    <div class="animate-spin rounded-full h-12 w-12 border-4 border-primary-500 border-t-transparent"></div>
  </div>
</template>
