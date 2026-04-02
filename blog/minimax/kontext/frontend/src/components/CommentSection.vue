<script setup lang="ts">
import { ref } from 'vue'
import { useCommentStore } from '@/stores/comment'
import CommentItem from './CommentItem.vue'
import { timeAgo } from '@/utils/time'

const props = defineProps<{
  articleId: number
}>()

const commentStore = useCommentStore()
const nickname = ref('')
const email = ref('')
const content = ref('')
const submitting = ref(false)

async function submitComment() {
  if (!nickname.value.trim() || !content.value.trim()) {
    alert('请填写昵称和评论内容')
    return
  }

  submitting.value = true
  try {
    await commentStore.createComment({
      article_id: props.articleId,
      nickname: nickname.value,
      email: email.value,
      content: content.value
    })
    nickname.value = ''
    email.value = ''
    content.value = ''
    alert('评论已提交，等待审核')
  } catch (e) {
    alert('评论提交失败')
  } finally {
    submitting.value = false
  }
}
</script>

<template>
  <div class="mt-12 bg-white rounded-xl shadow-md p-8">
    <h2 class="text-2xl font-bold mb-6">评论</h2>

    <div class="mb-8">
      <h3 class="text-lg font-semibold mb-4">发表评论</h3>
      <div class="grid md:grid-cols-2 gap-4 mb-4">
        <input
          v-model="nickname"
          type="text"
          placeholder="昵称 *"
          class="input"
        />
        <input
          v-model="email"
          type="email"
          placeholder="邮箱（选填）"
          class="input"
        />
      </div>
      <textarea
        v-model="content"
        placeholder="评论内容 *"
        rows="4"
        class="input mb-4"
      ></textarea>
      <button
        @click="submitComment"
        :disabled="submitting"
        class="btn btn-primary"
      >
        {{ submitting ? '提交中...' : '提交评论' }}
      </button>
    </div>

    <div v-if="commentStore.comments.length > 0">
      <h3 class="text-lg font-semibold mb-4">评论列表 ({{ commentStore.comments.length }})</h3>
      <div class="space-y-4">
        <CommentItem
          v-for="comment in commentStore.comments"
          :key="comment.id"
          :comment="comment"
        />
      </div>
    </div>
    <div v-else class="text-gray-500 text-center py-8">
      暂无评论
    </div>
  </div>
</template>
