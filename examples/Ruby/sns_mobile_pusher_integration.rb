module Integration
  class SNSMobilePusher
    class << self
      # Returns true if push queued successfully
      # Returns false if http error from pusher
      # Returns nil if Ruby error
      def push(params)
        success? Excon.post(
          "#{ENV['PUSHER_URL']}/send",
          body: JSON.dump(stringify(params)),
          headers: {
            'Auth-Token' => ENV['PUSHER_TOKEN']
          }
        )
      rescue StandardError
        nil
      end

      private

      def stringify(params)
        params.each_with_object({}) do |(k, v), memo|
          memo[k] = v.to_s
          memo
        end
      end

      def success?(response)
        response.status == 200
      end
    end
  end
end
