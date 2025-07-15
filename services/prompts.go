package services

// WorkoutPlanPrompt is the system prompt for generating workout plans
const WorkoutPlanPrompt = `You are a world-class fitness trainer, certified personal trainer, and exercise physiologist with 20+ years of experience. You have trained Olympic athletes, bodybuilders, and everyday people. You specialize in creating scientifically-backed, personalized workout plans that deliver results.
Your expertise includes:
- Exercise physiology and biomechanics
- Progressive overload and periodization
- Injury prevention and rehabilitation
- Nutrition and recovery optimization
- Sports psychology and motivation

Your task is to generate comprehensive, personalized workout plans in JSON format only. Always return valid JSON that matches the exact structure requested. Never include any text outside the JSON object.
Core principles you must follow:
- Evidence-based exercise selection
- Progressive overload for continuous improvement
- Proper exercise form and safety first
- Balanced muscle group targeting
- Appropriate rest periods and recovery
- Realistic weight recommendations based on fitness level
- Mix of compound and isolation exercises
- Consider user's available equipment and specific goals
- Periodization and variation for long-term success`

// WorkoutPlanTemplate is the template for generating workout plans
const WorkoutPlanTemplate = `Generate a world-class, personalized workout plan for the following user:
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

WORKOUT PLAN REQUIREMENTS:

STRUCTURE:
- Create 3-4 workout sessions per week based on fitness level
- Each session should have 4-8 exercises (depending on fitness level)
- Include warm-up and cool-down recommendations
- Balance push/pull movements across the week
- Progressive overload with realistic weight increases

EXERCISE SELECTION:
- Start with compound movements (squats, deadlifts, bench press, overhead press, rows)
- Follow with isolation exercises (curls, extensions, raises, flyes)
- Include core work in most sessions (planks, crunches, leg raises)
- Add cardio/conditioning where appropriate
- Use available equipment: %v

REPS AND SETS GUIDELINES:
- Beginner: 3 sets x 12-15 reps (focus on form)
- Intermediate: 4 sets x 8-12 reps (hypertrophy focus)
- Advanced: 4-5 sets x 6-8 reps (strength) or 8-12 reps (hypertrophy)

REST PERIODS:
- Compound movements: 2-3 minutes
- Isolation exercises: 60-90 seconds
- Supersets: 30-60 seconds between exercises

WEIGHT RECOMMENDATIONS:
- Beginner: Bodyweight or light weights, focus on form
- Intermediate: Moderate weights, 8-12 reps for hypertrophy
- Advanced: Heavier weights, 6-8 reps for strength, 8-12 for hypertrophy

SAFETY AND FORM:
- Always prioritize proper form over weight
- Include form cues and safety notes
- Consider user's experience level
- Provide progression guidelines

Please generate a comprehensive workout plan in the following JSON format:

{
  "id": 1,
  "name": "Professional Plan Name (e.g., 'Intermediate Strength Builder', 'Beginner Full Body Foundation')",
  "description": "Detailed description explaining the plan's scientific approach, expected results, and methodology. Include benefits, timeline, and what makes this plan effective for the user's specific situation.",
  "createdAt": "2024-01-15T00:00:00.000Z",
  "aiFeedbackCycle": 12,
  "planValidityPeriod": 28,
  "sessionsCompleted": 0,
  "planStartDate": "2024-01-15T00:00:00.000Z",
  "hasNewPlanSuggestion": false,
  "sessions": [
    {
      "id": "session_1",
      "name": "Descriptive Session Name (e.g., 'Upper Body Power', 'Lower Body Strength', 'Full Body Conditioning')",
      "note": "Comprehensive notes including: focus areas, form cues, breathing patterns, tempo recommendations, safety considerations, and session goals. Include specific instructions for each exercise type.",
      "exercises": [
        {
          "id": 1,
          "name": "Specific Exercise Name (e.g., 'Barbell Squat', 'Dumbbell Bench Press', 'Pull-up')",
          "sets": 3,
          "reps": 10,
          "weight": {"value": 100, "unit": "LB"},
          "type": "weight"
        }
      ]
    }
  ]
}

SPECIFIC INSTRUCTIONS FOR OPTIMAL RESULTS:

1. EQUIPMENT UTILIZATION: Use available equipment: %v
2. GOAL FOCUS: Tailor to user goals: %v
3. FITNESS LEVEL ADAPTATION: Adjust difficulty based on fitness level (%s)
4. WEIGHT UNITS: Use appropriate weight units (%s)
5. EXERCISE VARIATIONS: Include specific exercise variations based on available equipment
6. REALISTIC WEIGHTS: Provide realistic weight recommendations based on user's stats
7. FORM CUES: Include detailed form cues and safety notes
8. PROGRESSION: Consider user's experience level and provide clear progression guidelines
9. BALANCE: Ensure balanced push/pull movements throughout the week
10. MOBILITY: Include mobility and flexibility work where appropriate
11. RECOVERY: Consider rest days and recovery between sessions
12. PERSONALIZATION: Make exercises specific to user's goals and equipment

EXERCISE NAMING CONVENTIONS:
- Be specific: "Barbell Squat" not just "Squat"
- Include variations: "Dumbbell Bench Press", "Incline Barbell Press"
- Specify equipment: "Cable Row", "Lat Pulldown", "Smith Machine Squat"

WEIGHT RECOMMENDATIONS BY FITNESS LEVEL:
- Beginner: Start with bodyweight or very light weights
- Intermediate: Moderate weights based on user's stats
- Advanced: Heavier weights with proper progression

Return only the JSON object, no additional text or explanations. Ensure all exercise names are specific and professional.`
