import type { Meta, StoryObj } from '@storybook/react';
import { expect, userEvent, within } from 'storybook/test';

import Datepicker from './Datepicker';
import '../index.css';

const meta: Meta<typeof Datepicker> = {
  title: 'Components/Datepicker',
  component: Datepicker,
  parameters: {
    layout: 'centered',
    docs: {
      description: {
        component: 'A customizable datepicker component with calendar popup functionality.',
      },
    },
  },
  tags: ['autodocs'],
  argTypes: {
    // Since the component doesn't have props, we can add some documentation here
  },
};

export default meta;
type Story = StoryObj<typeof meta>;

// Default story
export const Default: Story = {
  parameters: {
    docs: {
      description: {
        story: 'The default datepicker with no date selected.',
      },
    },
  },
};

// Story with container to show positioning
export const WithContainer: Story = {
  parameters: {
    docs: {
      description: {
        story: 'Datepicker shown within a container to demonstrate positioning.',
      },
    },
  },
  decorators: [
    (Story) => (
      <div className="p-8 border border-dashed border-gray-300 rounded-lg min-h-[300px]">
        <h3 className="mb-4 text-gray-700">
          Select a date below:
        </h3>
        <Story />
      </div>
    ),
  ],
};

// Story with multiple datepickers
export const MultipleDatepickers: Story = {
  parameters: {
    docs: {
      description: {
        story: 'Multiple datepicker instances to show independent functionality.',
      },
    },
  },
  render: () => (
    <div className="flex gap-4 flex-col">
      <div>
        <div className="block mb-2 font-semibold">
          Start Date:
        </div>
        <Datepicker />
      </div>
      <div>
        <div className="block mb-2 font-semibold">
          End Date:
        </div>
        <Datepicker />
      </div>
    </div>
  ),
};

// Story with form context
export const InForm: Story = {
  parameters: {
    docs: {
      description: {
        story: 'Datepicker used within a form context.',
      },
    },
  },
  render: () => (
    <form className="p-6 border border-gray-200 rounded-lg bg-gray-50 max-w-sm">
      <h3 className="mb-4 text-gray-700">
        Event Registration
      </h3>
      
      <div className="mb-4">
        <label htmlFor="fullName" className="block mb-2 font-semibold">
          Full Name:
        </label>
        <input 
          id="fullName"
          type="text" 
          placeholder="Enter your name"
          className="w-full p-2 border border-gray-300 rounded-md"
        />
      </div>
      
      <div className="mb-4">
        <div className="block mb-2 font-semibold">
          Event Date:
        </div>
        <Datepicker />
      </div>
      
      <button 
        type="submit"
        className="bg-emerald-500 text-white py-2 px-4 border-0 rounded-md cursor-pointer hover:bg-emerald-600"
      >
        Register
      </button>
    </form>
  ),
};

// Story showing calendar interaction
export const CalendarInteraction: Story = {
  parameters: {
    docs: {
      description: {
        story: 'Demonstrates the calendar popup and date selection functionality. Click the input to open the calendar.',
      },
    },
  },
  decorators: [
    (Story) => (
      <div className="p-8 min-h-[400px] flex flex-col items-center">
        <p className="mb-4 text-gray-500 text-center max-w-xs">
          Click on the input field below to open the calendar and select a date.
          You can navigate between months using the arrow buttons.
        </p>
        <Story />
      </div>
    ),
  ],
};

// Interaction test: Open calendar
export const OpenCalendar: Story = {
  parameters: {
    docs: {
      description: {
        story: 'Tests opening the calendar popup by clicking the input field.',
      },
    },
  },
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);
    const input = canvas.getByPlaceholderText('Select date');
    
    // Test that calendar is initially closed
    expect(canvas.queryByText(/January|February|March|April|May|June|July|August|September|October|November|December/)).not.toBeInTheDocument();
    
    // Click to open calendar
    await userEvent.click(input);
    
    // Test that calendar is now open
    expect(canvas.getByText(/January|February|March|April|May|June|July|August|September|October|November|December/)).toBeInTheDocument();
  },
};

// Interaction test: Select a date
export const SelectDate: Story = {
  parameters: {
    docs: {
      description: {
        story: 'Tests selecting a date from the calendar.',
      },
    },
  },
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);
    const input = canvas.getByPlaceholderText('Select date');
    
    // Open calendar
    await userEvent.click(input);
    
    // Find and click on day 15 (assuming it's in current month)
    const day15 = canvas.getByRole('button', { name: '15' });
    await userEvent.click(day15);
    
    // Test that calendar closes after selection
    expect(canvas.queryByText(/January|February|March|April|May|June|July|August|September|October|November|December/)).not.toBeInTheDocument();
    
    // Test that input field shows selected date
    expect((input as HTMLInputElement).value).not.toBe('');
  },
};

// Interaction test: Navigate months
export const NavigateMonths: Story = {
  parameters: {
    docs: {
      description: {
        story: 'Tests navigating between months using arrow buttons.',
      },
    },
  },
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);
    const input = canvas.getByPlaceholderText('Select date');
    
    // Open calendar
    await userEvent.click(input);
    
    // Get current month name
    const monthHeader = canvas.getByText(/January|February|March|April|May|June|July|August|September|October|November|December/);
    const currentMonth = monthHeader.textContent;
    
    // Click next month button
    const nextButton = canvas.getByRole('button', { name: '›' });
    await userEvent.click(nextButton);
    
    // Test that month changed
    const newMonthHeader = canvas.getByText(/January|February|March|April|May|June|July|August|September|October|November|December/);
    expect(newMonthHeader.textContent).not.toBe(currentMonth);
    
    // Click previous month button
    const prevButton = canvas.getByRole('button', { name: '‹' });
    await userEvent.click(prevButton);
    
    // Test that we're back to original month
    const finalMonthHeader = canvas.getByText(/January|February|March|April|May|June|July|August|September|October|November|December/);
    expect(finalMonthHeader.textContent).toBe(currentMonth);
  },
};

// Interaction test: Close calendar by clicking outside
export const CloseCalendarOutside: Story = {
  parameters: {
    docs: {
      description: {
        story: 'Tests that the calendar closes when clicking outside of it.',
      },
    },
  },
  decorators: [
    (Story) => (
      <div className="p-8 min-h-[400px]">
        <div className="mb-4">
          <Story />
        </div>
        <div data-testid="outside-area" className="h-32 bg-gray-100 rounded p-4">
          Click here to test outside click
        </div>
      </div>
    ),
  ],
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);
    const input = canvas.getByPlaceholderText('Select date');
    
    // Open calendar
    await userEvent.click(input);
    
    // Verify calendar is open
    expect(canvas.getByText(/January|February|March|April|May|June|July|August|September|October|November|December/)).toBeInTheDocument();
    
    // Click outside area
    const outsideArea = canvas.getByTestId('outside-area');
    await userEvent.click(outsideArea);
    
    // Test that calendar closes
    expect(canvas.queryByText(/January|February|March|April|May|June|July|August|September|October|November|December/)).not.toBeInTheDocument();
  },
};

// Interaction test: Multiple interactions
export const CompleteInteraction: Story = {
  parameters: {
    docs: {
      description: {
        story: 'Tests a complete user interaction flow: open calendar, navigate months, select date.',
      },
    },
  },
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement);
    const input = canvas.getByPlaceholderText('Select date');
    
    // 1. Open calendar
    await userEvent.click(input);
    expect(canvas.getByText(/January|February|March|April|May|June|July|August|September|October|November|December/)).toBeInTheDocument();
    
    // 2. Navigate to next month
    const nextButton = canvas.getByRole('button', { name: '›' });
    await userEvent.click(nextButton);
    
    // 3. Select a date (day 10)
    const day10 = canvas.getByRole('button', { name: '10' });
    await userEvent.click(day10);
    
    // 4. Verify calendar closes and date is selected
    expect(canvas.queryByText(/January|February|March|April|May|June|July|August|September|October|November|December/)).not.toBeInTheDocument();
    expect((input as HTMLInputElement).value).not.toBe('');
    
    // 5. Reopen calendar to verify selected date is highlighted
    await userEvent.click(input);
    
    // The selected day should have a different style (bg-web-green-500)
    const selectedDay = canvas.getByRole('button', { name: '10' });
    expect(selectedDay).toHaveClass('bg-web-green-500');
  },
};