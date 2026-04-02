<script setup lang="ts">
import type { Comment } from '@/types'
import { timeAgo } from '@/utils/time'

defineProps<{
  comment: Comment
}>()
</script>

<template>
  <div class="border-b border-gray-100 pb-4">
    <div class="flex items-center gap-3 mb-2">
      <div class="w-8 h-8 bg-primary-100 rounded-full flex items-center justify-center text-primary-700 font-bold">
        {{ comment.nickname.charAt(0).toUpperCase() }}
      </div>
      <div>
        <span class="font-semibold">{{ comment.nickname }}</span>
        <span class="text-gray-400 text-sm ml-2">{{ timeAgo(comment.created_at) }}</span>
      </div>
    </div>
    <p class="text-gray-700 pl-11">{{ comment.content }}</p>
  </div>
</template>
