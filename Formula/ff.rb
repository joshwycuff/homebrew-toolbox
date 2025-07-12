class Ff < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/ff-654f46bc058ea34cb4bf1373b439b902cf699b1f/ff.tar.gz"
  version "654f46bc058ea34cb4bf1373b439b902cf699b1f"

  depends_on "rg"
  depends_on "bat"

  def install
    bin.install "ff"
    man1.install Dir["*.1"]
    bash_completion.install "ff.bash" => "ff"
    zsh_completion.install "ff.zsh" => "_ff"
    fish_completion.install "ff.fish"
  end
end
