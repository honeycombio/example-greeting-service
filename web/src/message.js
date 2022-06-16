import { trace, context } from '@opentelemetry/api';
import { tracer } from "./tracing-http";

const messages = [
  'how are you?',
  'how are you doing?',
  "what's good?",
  "what's up?",
  'how do you do?',
  'sup?',
  'good day to you',
  'how are things?',
  'howzit?',
  'woohoo',
];

export function determineMessage() {
  const message = messages[Math.floor(Math.random() * messages.length)];
  return messages[Math.floor(Math.random() * messages.length)];
}