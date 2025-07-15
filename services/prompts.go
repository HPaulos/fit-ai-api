package services

// WorkoutPlanPrompt is the system prompt for generating workout plans
const WorkoutPlanPrompt = `You are an expert fitness trainer and nutritionist with 15+ years of experience. You specialize in creating personalized workout plans that are safe, effective, and tailored to individual needs. 

Your task is to generate detailed workout plans in JSON format only. Always return valid JSON that matches the exact structure requested. Never include any text outside the JSON object.

Key principles to follow:
- Progressive overload for strength gains
- Proper exercise form and safety
- Balanced muscle group targeting
- Appropriate rest periods
- Realistic weight recommendations based on fitness level
- Mix of compound and isolation exercises
- Consider user's available equipment and goals`

// WorkoutPlanTemplate is the template for generating workout plans
const WorkoutPlanTemplate = `Generate a comprehensive, personalized workout plan for the following user:

USER PROFILE:
- Name: %s
- Age: %s (calculated from date of birth)
- Gender: %s
- Fitness Level: %s
- Activity Level: %s
- Height: %.1f %s
- Weight: %.1f %s
- Goals: %v
- Available Equipment: %v
- Location: %s
- Units: %s

CURRENT STATS:
- Total Workouts: %d
- Current Streak: %d
- Longest Streak: %d
- Total Time: %d minutes
- Total Volume: %d

REQUIREMENTS:
Create a workout plan that includes:
1. 3-4 workout sessions per week based on fitness level
2. Each session should have 4-6 exercises
3. Mix of compound and isolation movements
4. Progressive overload principles
5. Appropriate rest days between muscle groups
6. Warm-up and cool-down recommendations

EXERCISE GUIDELINES:
- Compound exercises first (squats, deadlifts, bench press, etc.)
- Isolation exercises second (curls, extensions, raises, etc.)
- Core exercises included in most sessions
- Cardio recommendations where appropriate
- Rest periods: 60-90 seconds for hypertrophy, 2-3 minutes for strength

WEIGHT RECOMMENDATIONS:
- Beginner: Focus on form, bodyweight or light weights
- Intermediate: Moderate weights, 8-12 reps for hypertrophy
- Advanced: Heavier weights, 6-8 reps for strength, 8-12 for hypertrophy

Please generate a complete workout plan in the following JSON format:

{
  "id": 1,
  "name": "Descriptive Plan Name",
  "description": "Detailed description explaining the plan's focus, benefits, and approach",
  "createdAt": "2024-01-15T00:00:00.000Z",
  "aiFeedbackCycle": 12,
  "planValidityPeriod": 28,
  "sessionsCompleted": 0,
  "planStartDate": "2024-01-15T00:00:00.000Z",
  "hasNewPlanSuggestion": false,
  "sessions": [
    {
      "id": "session_1",
      "name": "Descriptive Session Name",
      "note": "Detailed notes about focus, form cues, and session goals",
      "exercises": [
        {
          "id": 1,
          "name": "Specific Exercise Name",
          "sets": 3,
          "reps": 10,
          "weight": {"value": 100, "unit": "LB"},
          "type": "weight"
        }
      ]
    }
  ]
}

SPECIFIC INSTRUCTIONS:
1. Use available equipment: %v
2. Focus on user goals: %v
3. Adjust difficulty based on fitness level (%s)
4. Use appropriate weight units (%s)
5. Include specific exercise variations based on equipment
6. Provide realistic weight recommendations
7. Include form cues and safety notes
8. Consider user's experience level and progression
9. Balance push/pull movements
10. Include mobility and flexibility work

Return only the JSON object, no additional text or explanations.`
